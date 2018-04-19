'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var DocumentClickHandler = function DocumentClickHandler(WrappedComponent) {
  var _class, _temp;

  return _temp = _class = function (_React$PureComponent) {
    _inherits(_class, _React$PureComponent);

    function _class(props) {
      _classCallCheck(this, _class);

      var _this = _possibleConstructorReturn(this, (_class.__proto__ || Object.getPrototypeOf(_class)).call(this, props));

      _this.notifyClose = _this.notifyClose.bind(_this);
      return _this;
    }

    _createClass(_class, [{
      key: 'componentDidMount',
      value: function componentDidMount() {
        if (this.props.open) {
          document.addEventListener('click', this.notifyClose);
        }
      }
    }, {
      key: 'componentWillUnmount',
      value: function componentWillUnmount() {
        document.removeEventListener('click', this.notifyClose);
      }
    }, {
      key: 'componentWillReceiveProps',
      value: function componentWillReceiveProps(_ref) {
        var nextOpen = _ref.open;
        var open = this.props.open;

        if (open && !nextOpen) {
          document.removeEventListener('click', this.notifyClose);
        } else if (!open && nextOpen) {
          document.addEventListener('click', this.notifyClose);
        }
      }
    }, {
      key: 'notifyClose',
      value: function notifyClose() {
        if (this.props.onClose) {
          this.props.onClose();
        }
      }
    }, {
      key: 'render',
      value: function render() {
        var _props = this.props,
            onClose = _props.onClose,
            otherProps = _objectWithoutProperties(_props, ['onClose']);

        return _react2.default.createElement(WrappedComponent, otherProps);
      }
    }]);

    return _class;
  }(_react2.default.PureComponent), _class.propTypes = {
    open: _propTypes2.default.bool,
    onClose: _propTypes2.default.func
  }, _temp;
};
exports.default = DocumentClickHandler;